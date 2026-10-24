/**
 * Storage abstraction for default models.
 * 
 * This module provides functions to get and save default models.
 * For the purpose of this implementation, it uses an in-memory store.
 */

import { DefaultModel } from './defaultModelRefresher';

let inMemoryDefaultModels: DefaultModel[] = [];

/**
 * Gets the current default models.
 */
export async function getDefaultModels(): Promise<DefaultModel[]> {
  // Simulate async storage read
  return inMemoryDefaultModels;
}

/**
 * Saves the given default models.
 * @param models - Array of default models to save.
 */
export async function saveDefaultModels(models: DefaultModel[]): Promise<void> {
  // Simulate async storage write
  inMemoryDefaultModels = models;
}
