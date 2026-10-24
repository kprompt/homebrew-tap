import fs from 'fs';
import path from 'path';

export interface Model {
  modelName: string;
  version: string;
  description: string;
  parameters: Record<string, any>;
  metadata: Record<string, any>;
}

const DEFAULT_MODEL_PATH = path.resolve(__dirname, '../../models/default_model.json');

let cachedModel: Model | null = null;

/**
 * Loads the default model from disk.
 * Caches the model after first load.
 */
export function loadDefaultModel(): Model {
  if (cachedModel) {
    return cachedModel;
  }
  const raw = fs.readFileSync(DEFAULT_MODEL_PATH, 'utf-8');
  const model = JSON.parse(raw) as Model;
  cachedModel = model;
  return model;
}
