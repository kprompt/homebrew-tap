/**
 * Module to refresh stale default models.
 * This module checks if the default models are stale based on a timestamp,
 * and refreshes them if needed.
 */

import fs from 'fs/promises';
import path from 'path';

const MODEL_DIR = path.resolve(__dirname, '../../models');
const DEFAULT_MODEL_FILE = path.join(MODEL_DIR, 'defaultModel.json');
const STALE_THRESHOLD_MS = 1000 * 60 * 60 * 24 * 7; // 7 days

export interface DefaultModel {
  version: string;
  data: any;
  lastUpdated: number; // timestamp in ms
}

/**
 * Loads the default model from disk.
 */
export async function loadDefaultModel(): Promise<DefaultModel | null> {
  try {
    const content = await fs.readFile(DEFAULT_MODEL_FILE, 'utf-8');
    return JSON.parse(content) as DefaultModel;
  } catch (err) {
    // File not found or invalid JSON
    return null;
  }
}

/**
 * Saves the default model to disk.
 * @param model The model to save.
 */
export async function saveDefaultModel(model: DefaultModel): Promise<void> {
  await fs.mkdir(MODEL_DIR, { recursive: true });
  await fs.writeFile(DEFAULT_MODEL_FILE, JSON.stringify(model, null, 2), 'utf-8');
}

/**
 * Checks if the default model is stale.
 * @param model The model to check.
 * @returns True if stale, false otherwise.
 */
export function isModelStale(model: DefaultModel): boolean {
  const now = Date.now();
  return now - model.lastUpdated > STALE_THRESHOLD_MS;
}

/**
 * Fetches a fresh default model from the provider.
 * This simulates an async fetch operation.
 */
export async function fetchFreshDefaultModel(): Promise<DefaultModel> {
  // Simulate fetching fresh model data from a provider
  return {
    version: `v${Date.now()}`,
    data: { refreshed: true, timestamp: new Date().toISOString() },
    lastUpdated: Date.now(),
  };
}

/**
 * Refreshes the default model if it is stale or missing.
 * @returns The refreshed or existing model.
 */
export async function refreshDefaultModelIfStale(): Promise<DefaultModel> {
  const currentModel = await loadDefaultModel();
  if (!currentModel || isModelStale(currentModel)) {
    const freshModel = await fetchFreshDefaultModel();
    await saveDefaultModel(freshModel);
    return freshModel;
  }
  return currentModel;
}
