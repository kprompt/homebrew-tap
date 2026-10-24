import { refreshDefaultModels } from '../../src/models/defaultModelRefresher';
import * as modelStorage from '../../src/models/modelStorage';
import * as modelFetcher from '../../src/models/modelFetcher';

jest.useFakeTimers('modern');

describe('refreshDefaultModels', () => {
  const now = Date.now();

  beforeEach(() => {
    jest.setSystemTime(now);
    jest.clearAllMocks();
  });

  it('returns current models if not stale', async () => {
    const currentModels = [
      { id: 'modelA', data: { name: 'A' }, lastUpdated: now - 1000 },
      { id: 'modelB', data: { name: 'B' }, lastUpdated: now - 1000 },
    ];
    jest.spyOn(modelStorage, 'getDefaultModels').mockResolvedValue(currentModels);
    const saveSpy = jest.spyOn(modelStorage, 'saveDefaultModels');

    const result = await refreshDefaultModels();

    expect(result).toEqual(currentModels);
    expect(saveSpy).not.toHaveBeenCalled();
  });

  it('refreshes models if stale', async () => {
    const staleModels = [
      { id: 'modelA', data: { name: 'A' }, lastUpdated: now - 25 * 60 * 60 * 1000 },
    ];
    const fetchedModels = [
      { id: 'modelA', data: { name: 'Model A', version: '1.1' } },
      { id: 'modelB', data: { name: 'Model B', version: '2.0' } },
    ];

    jest.spyOn(modelStorage, 'getDefaultModels').mockResolvedValue(staleModels);
    jest.spyOn(modelFetcher, 'fetchLatestDefaultModels').mockResolvedValue(fetchedModels);
    const saveSpy = jest.spyOn(modelStorage, 'saveDefaultModels').mockResolvedValue();

    const result = await refreshDefaultModels();

    expect(saveSpy).toHaveBeenCalledTimes(1);
    expect(result).toHaveLength(2);
    expect(result[0]).toMatchObject({ id: 'modelA', data: { name: 'Model A', version: '1.1' } });
    expect(result[0].lastUpdated).toBe(now);
  });

  it('refreshes models if no current models', async () => {
    jest.spyOn(modelStorage, 'getDefaultModels').mockResolvedValue([]);
    const fetchedModels = [
      { id: 'modelA', data: { name: 'Model A', version: '1.1' } },
    ];
    jest.spyOn(modelFetcher, 'fetchLatestDefaultModels').mockResolvedValue(fetchedModels);
    const saveSpy = jest.spyOn(modelStorage, 'saveDefaultModels').mockResolvedValue();

    const result = await refreshDefaultModels();

    expect(saveSpy).toHaveBeenCalled();
    expect(result).toHaveLength(1);
    expect(result[0].lastUpdated).toBe(now);
  });
});
