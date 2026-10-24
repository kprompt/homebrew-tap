const { getDefaultModels, refreshDefaultModels } = require('../src/models/defaultModels');

describe('defaultModels module', () => {
  jest.useFakeTimers('modern');

  beforeEach(() => {
    jest.setSystemTime(Date.now());
  });

  afterAll(() => {
    jest.useRealTimers();
  });

  test('getDefaultModels returns current models without refresh if not stale', () => {
    const models1 = getDefaultModels();
    const models2 = getDefaultModels();
    expect(models2).toEqual(models1);
  });

  test('getDefaultModels refreshes stale models', () => {
    const models1 = getDefaultModels();

    // Advance time beyond TTL to make models stale
    jest.advanceTimersByTime(1000 * 60 * 60 * 24 + 1);

    const models2 = getDefaultModels();

    // Models should be refreshed with incremented version and refreshed flag
    Object.keys(models1).forEach((modelName) => {
      expect(models2[modelName].version).toBe(models1[modelName].version + 1);
      expect(models2[modelName].data.refreshed).toBe(true);
      expect(models2[modelName].lastUpdated).toBeGreaterThan(models1[modelName].lastUpdated);
    });
  });

  test('refreshDefaultModels forces refresh regardless of staleness', () => {
    const models1 = getDefaultModels();

    const models2 = refreshDefaultModels();

    Object.keys(models1).forEach((modelName) => {
      expect(models2[modelName].version).toBe(models1[modelName].version + 1);
      expect(models2[modelName].data.refreshed).toBe(true);
      expect(models2[modelName].lastUpdated).toBeGreaterThan(models1[modelName].lastUpdated);
    });
  });
});
