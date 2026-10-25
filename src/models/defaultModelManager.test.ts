import fs from 'fs';
import path from 'path';
import { defaultModelManager } from './defaultModelManager';

jest.useFakeTimers();

const DEFAULT_MODEL_PATH = path.resolve(__dirname, '../../models/defaultModel.json');

describe('DefaultModelManager', () => {
  const sampleModel = {
    version: '1.0.0',
    data: { foo: 'bar' },
    lastUpdated: Date.now(),
  };

  beforeEach(() => {
    jest.resetModules();
    jest.clearAllTimers();
    // Mock fs.readFileSync to return sampleModel JSON string
    jest.spyOn(fs, 'readFileSync').mockImplementation(() => JSON.stringify(sampleModel));
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('loads the default model on initialization', () => {
    const model = defaultModelManager.getModel();
    expect(model).not.toBeNull();
    expect(model?.version).toBe(sampleModel.version);
  });

  it('refreshes the model if stale', () => {
    const manager = defaultModelManager as any;
    // Force lastLoadTime to be stale
    manager.lastLoadTime = Date.now() - (1000 * 60 * 60 * 25); // 25 hours ago

    // Spy on loadModel
    const loadSpy = jest.spyOn(manager, 'loadModel');

    manager.refreshIfStale();

    expect(loadSpy).toHaveBeenCalled();
  });

  it('does not refresh if model is fresh', () => {
    const manager = defaultModelManager as any;
    // Force lastLoadTime to be recent
    manager.lastLoadTime = Date.now();

    const loadSpy = jest.spyOn(manager, 'loadModel');

    manager.refreshIfStale();

    expect(loadSpy).not.toHaveBeenCalled();
  });

  it('periodically refreshes the model', () => {
    const manager = defaultModelManager as any;
    const loadSpy = jest.spyOn(manager, 'loadModel');

    // Advance time by 24 hours + 1 ms
    jest.advanceTimersByTime(1000 * 60 * 60 * 24 + 1);

    // The interval should trigger refreshIfStale which calls loadModel if stale
    expect(loadSpy).toHaveBeenCalled();
  });
});
