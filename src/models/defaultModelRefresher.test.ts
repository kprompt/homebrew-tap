import { refreshStaleDefaultModels } from './defaultModelRefresher';
import * as modelStore from './modelStore';

describe('refreshStaleDefaultModels', () => {
  beforeEach(() => {
    // Reset models to known state before each test
    jest.spyOn(modelStore, 'getDefaultModels').mockResolvedValue([
      {
        id: 'model-1',
        name: 'Default Model A',
        lastRefreshed: Date.now() - 1000 * 60 * 60 * 25, // 25 hours ago
        isDefault: true,
        data: { version: 1 },
      },
      {
        id: 'model-2',
        name: 'Default Model B',
        lastRefreshed: Date.now() - 1000 * 60 * 60 * 10, // 10 hours ago
        isDefault: true,
        data: { version: 1 },
      },
      {
        id: 'model-3',
        name: 'User Model C',
        lastRefreshed: Date.now() - 1000 * 60 * 60 * 30, // 30 hours ago
        isDefault: false,
        data: { version: 1 },
      },
    ]);
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('refreshes only stale default models', async () => {
    const updateModelMock = jest.spyOn(modelStore, 'updateModel').mockResolvedValue();

    const staleThresholdMs = 1000 * 60 * 60 * 24; // 24 hours

    const refreshedIds = await refreshStaleDefaultModels({ staleThresholdMs });

    expect(refreshedIds).toContain('model-1');
    expect(refreshedIds).not.toContain('model-2');
    expect(refreshedIds).not.toContain('model-3'); // not default

    expect(updateModelMock).toHaveBeenCalledTimes(1);
    expect(updateModelMock.mock.calls[0][0].id).toBe('model-1');
    expect(updateModelMock.mock.calls[0][0].lastRefreshed).toBeGreaterThan(Date.now() - 1000);

    // The refreshed data should include refreshedAt timestamp
    expect(updateModelMock.mock.calls[0][0].data.refreshedAt).toBeDefined();
  });

  it('returns empty array if no stale models', async () => {
    jest.spyOn(modelStore, 'getDefaultModels').mockResolvedValue([
      {
        id: 'model-4',
        name: 'Default Model D',
        lastRefreshed: Date.now() - 1000 * 60 * 60 * 1, // 1 hour ago
        isDefault: true,
        data: { version: 1 },
      },
    ]);
    const updateModelMock = jest.spyOn(modelStore, 'updateModel').mockResolvedValue();

    const refreshedIds = await refreshStaleDefaultModels({ staleThresholdMs: 1000 * 60 * 60 * 2 }); // 2 hours

    expect(refreshedIds).toEqual([]);
    expect(updateModelMock).not.toHaveBeenCalled();
  });

  it('handles updateModel failure gracefully', async () => {
    jest.spyOn(modelStore, 'updateModel').mockRejectedValue(new Error('DB error'));
    const consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
    const consoleInfoSpy = jest.spyOn(console, 'info').mockImplementation(() => {});

    const refreshedIds = await refreshStaleDefaultModels({ staleThresholdMs: 1000 * 60 * 60 * 24 });

    expect(refreshedIds).toEqual([]);
    expect(consoleErrorSpy).toHaveBeenCalled();

    consoleErrorSpy.mockRestore();
    consoleInfoSpy.mockRestore();
  });
});
