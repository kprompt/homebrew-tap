import {
  BaseProvider,
  ProviderOptions,
  ProviderResponse,
  ProviderRequest,
} from '../types/providerTypes';
import fetch from 'node-fetch';

interface AmazonBedrockOptions extends ProviderOptions {
  endpoint: string;
  apiKey: string;
  modelId: string;
}

export class AmazonBedrockProvider extends BaseProvider {
  private endpoint: string;
  private apiKey: string;
  private modelId: string;

  constructor(options: AmazonBedrockOptions) {
    super('amazon-bedrock', options);
    this.endpoint = options.endpoint;
    this.apiKey = options.apiKey;
    this.modelId = options.modelId;
  }

  async generate(request: ProviderRequest): Promise<ProviderResponse> {
    const url = `${this.endpoint}/models/${this.modelId}/invoke`;

    const body = {
      input: request.prompt,
      parameters: request.parameters || {},
    };

    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'x-api-key': this.apiKey,
      },
      body: JSON.stringify(body),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Amazon Bedrock API error: ${response.status} ${response.statusText} - ${errorText}`
      );
    }

    const json = await response.json();

    if (!json || !json.generatedText) {
      throw new Error('Amazon Bedrock API response missing generatedText');
    }

    return {
      text: json.generatedText,
      raw: json,
    };
  }
}
