import { loadDefaultModel } from '../src/models/modelLoader';
import path from 'path';
import fs from 'fs';

describe('loadDefaultModel', () => {
  const modelPath = path.resolve(__dirname, '../models/default_model.json');

  it('loads the default model correctly', () => {
    const model = loadDefaultModel();
    expect(model).toHaveProperty('modelName', 'default-model');
    expect(model).toHaveProperty('version', '1.1.0');
    expect(model.parameters).toHaveProperty('threshold', 0.75);
    expect(model.parameters).toHaveProperty('maxIterations', 1000);
    expect(model.parameters).toHaveProperty('learningRate', 0.01);
  });

  it('caches the model after first load', () => {
    const firstLoad = loadDefaultModel();
    const secondLoad = loadDefaultModel();
    expect(firstLoad).toBe(secondLoad);
  });

  it('model file exists and is valid JSON', () => {
    expect(fs.existsSync(modelPath)).toBe(true);
    const raw = fs.readFileSync(modelPath, 'utf-8');
    expect(() => JSON.parse(raw)).not.toThrow();
  });
});
