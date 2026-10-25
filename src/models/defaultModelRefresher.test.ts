import { refreshStaleDefaultModels } from './defaultModelRefresher';
import * as modelStorage from './modelStorage';
import { Logger } from './logger';

describe('refreshStaleDefaultModels', () => {
  const logger: Logger = {
    info: jest.fn(),
    warn: jest.fn(),
    error: jest.fn(),
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('refreshes models that are stale and skips fresh ones', async () => {
    const now = Date.now();
    const staleModel = { id: 'stale-model' };
    const freshModel = { id: 'fresh-model' };
    const noTimestampModel = { id: 'no-timestamp-model' };

    jest.spyOn(modelStorage, 'getDefaultModels').mockResolvedValue([
      staleModel,
      freshModel,
      noTimestampModel,
    ]);
    jest.spyOn(modelStorage, 'getModelLastUpdated').mockImplementation(async (id: string) => {
      if (id === 'stale-model') return now - 1000 * 60 * 60 * 24 * 8; // 8 days ago
      if (id === 'fresh-model') return now - 1000 * 60 * 60 * 24 * 2; // 2 days ago
      if (id === 'no-timestamp-model') return null;
      return null;
    });
    const reloadSpy = jest.spyOn(modelStorage, 'reloadModel').mockResolvedValue();

    const result = await refreshStaleDefaultModels(logger);

    expect(reloadSpy).toHaveBeenCalledTimes(2);
    expect(reloadSpy).toHaveBeenCalledWith('stale-model');
    expect(reloadSpy).toHaveBeenCalledWith('no-timestamp-model');
    expect(result.refreshedModels).toEqual(expect.arrayContaining(['stale-model', 'no-timestamp-model']));
    expect(result.skippedModels).toEqual(['fresh-model']);
    expect(logger.warn).toHaveBeenCalledWith(expect.stringContaining('no last updated timestamp'));
    expect(logger.info).toHaveBeenCalledWith(expect.stringContaining('Refreshed 2 stale default models.'));
  });

  it('handles empty default models gracefully', async () => {
    jest.spyOn(modelStorage, 'getDefaultModels').mockResolvedValue([]);
    const reloadSpy = jest.spyOn(modelStorage, 'reloadModel').mockResolvedValue();

    const result = await refreshStaleDefaultModels(logger);

    expect(reloadSpy).not.toHaveBeenCalled();
    expect(result.refreshedModels).toEqual([]);
    expect(result.skippedModels).toEqual([]);
    expect(logger.info).toHaveBeenCalledWith('Refreshed 0 stale default models.');
  });
});
