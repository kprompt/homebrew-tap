/**
 * Model storage and retrieval utilities.
 * 
 * This is a simple in-memory store for demonstration.
 * Replace with real DB or persistent store in production.
 */

import { Model } from './defaultModelRefresher';

const models: Model[] = [
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
];

/**
 * Returns all default models.
 */
export async function getDefaultModels(): Promise<Model[]> {
  // Simulate async DB call
  return models.filter((m) => m.isDefault);
}

/**
 * Updates a model in the store.
 * @param updatedModel 
 */
export async function updateModel(updatedModel: Model): Promise<void> {
  const idx = models.findIndex((m) => m.id === updatedModel.id);
  if (idx === -1) {
    throw new Error(`Model with id ${updatedModel.id} not found`);
  }
  models[idx] = updatedModel;
}
