/**
 * Mocked model storage interface.
 * 
 * In a real implementation, these would interact with a database or filesystem.
 */

export interface Model {
  id: string;
  // other model metadata
}

/**
 * Returns all default models.
 */
export async function getDefaultModels(): Promise<Model[]> {
  // Placeholder: In real code, fetch from DB or config
  return [];
}

/**
 * Returns the last updated timestamp (ms since epoch) for a model by id.
 * Returns null if no timestamp is available.
 */
export async function getModelLastUpdated(modelId: string): Promise<number | null> {
  // Placeholder: In real code, fetch from DB or metadata store
  return null;
}

/**
 * Reloads or rebuilds the model by id.
 */
export async function reloadModel(modelId: string): Promise<void> {
  // Placeholder: In real code, trigger model rebuild or reload
  return;
}
