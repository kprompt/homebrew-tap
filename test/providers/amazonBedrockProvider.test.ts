import { AmazonBedrockProvider } from '../../src/providers/amazonBedrockProvider';
import fetchMock from 'jest-fetch-mock';

fetchMock.enableMocks();

describe('AmazonBedrockProvider', () => {
  const endpoint = 'https://bedrock.us-east-1.amazonaws.com';
  const apiKey = 'test-api-key';
  const modelId = 'test-model';

  beforeEach(() => {
    fetchMock.resetMocks();
  });

  it('should send correct request and parse response', async () => {
    const prompt = 'Hello, world!';
    const generatedText = 'Hello from Bedrock!';
    fetchMock.mockResponseOnce(
      JSON.stringify({ generatedText }),
      { status: 200 }
    );

    const provider = new AmazonBedrockProvider({ endpoint, apiKey, modelId });
    const response = await provider.generate({ prompt });

    expect(fetchMock).toHaveBeenCalledTimes(1);
    const fetchCall = fetchMock.mock.calls[0];
    expect(fetchCall[0]).toBe(`${endpoint}/models/${modelId}/invoke`);
    const fetchOptions = fetchCall[1];
    expect(fetchOptions?.method).toBe('POST');
    expect(fetchOptions?.headers).toMatchObject({
      'Content-Type': 'application/json',
      'x-api-key': apiKey,
    });
    const body = JSON.parse(fetchOptions?.body as string);
    expect(body.input).toBe(prompt);

    expect(response.text).toBe(generatedText);
    expect(response.raw).toEqual({ generatedText });
  });

  it('should throw on non-200 response', async () => {
    fetchMock.mockResponseOnce('Unauthorized', { status: 401, statusText: 'Unauthorized' });

    const provider = new AmazonBedrockProvider({ endpoint, apiKey, modelId });
    await expect(provider.generate({ prompt: 'test' })).rejects.toThrow(
      /Amazon Bedrock API error: 401 Unauthorized/
    );
  });

  it('should throw if generatedText missing in response', async () => {
    fetchMock.mockResponseOnce(JSON.stringify({}), { status: 200 });

    const provider = new AmazonBedrockProvider({ endpoint, apiKey, modelId });
    await expect(provider.generate({ prompt: 'test' })).rejects.toThrow(
      /Amazon Bedrock API response missing generatedText/
    );
  });
});
