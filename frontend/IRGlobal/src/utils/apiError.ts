/**
 * Extrai a mensagem de erro do backend ou retorna erro padrão
 * @param response - Response da fetch API
 * @param defaultMessage - Mensagem padrão caso não consiga extrair do backend
 * @returns Promise<never> - Sempre lança um erro
 */
export async function handleApiError(
  response: Response,
  defaultMessage: string
): Promise<never> {
  // Verificar se o body já foi consumido
  if (response.bodyUsed) {
    throw new Error(
      `${defaultMessage}: ${response.status} ${response.statusText}`
    );
  }

  try {
    // Clonar a response para poder ler o body múltiplas vezes
    const responseClone = response.clone();
    const errorData = await responseClone.json();

    const backendError =
      errorData.error || `${response.status} ${response.statusText}`;
    // Lançar o erro com a mensagem do backend
    throw new Error(`${defaultMessage}: ${backendError}`);
  } catch (parseError) {
    // Verificar se o erro é da nossa própria exceção (que queremos propagar)
    if (
      parseError instanceof Error &&
      parseError.message.startsWith(defaultMessage)
    ) {
      throw parseError; // Re-lançar nossa própria exceção
    }

    // Tentar ler como texto se JSON falhar e body não foi usado
    if (!response.bodyUsed) {
      try {
        const responseText = await response.text();

        // Se o texto contém JSON, tentar extrair o erro
        if (responseText.includes('"error"')) {
          const match = responseText.match(/"error"\s*:\s*"([^"]+)"/);
          if (match) {
            throw new Error(`${defaultMessage}: ${match[1]}`);
          }
        }
      } catch (textError) {
        // Ignorar erro de leitura de texto
      }
    }

    // Se não conseguir fazer parse do JSON, usar erro padrão
    throw new Error(
      `${defaultMessage}: ${response.status} ${response.statusText}`
    );
  }
}

/**
 * Extrai apenas a mensagem de erro do backend
 * @param response - Response da fetch API
 * @returns Promise<string> - Mensagem de erro
 */
export async function extractErrorMessage(response: Response): Promise<string> {
  // Verificar se o body já foi consumido
  if (response.bodyUsed) {
    return `${response.status} ${response.statusText}`;
  }

  try {
    // Clonar a response para poder ler o body múltiplas vezes
    const responseClone = response.clone();
    const errorData = await responseClone.json();

    return errorData.error || `${response.status} ${response.statusText}`;
  } catch (parseError) {
    // Tentar ler como texto se JSON falhar
    if (!response.bodyUsed) {
      try {
        const responseText = await response.text();

        // Se o texto contém JSON, tentar extrair o erro
        if (responseText.includes('"error"')) {
          const match = responseText.match(/"error"\s*:\s*"([^"]+)"/);
          if (match) {
            return match[1];
          }
        }
      } catch (textError) {
        // Ignorar erro de leitura de texto
      }
    }

    return `${response.status} ${response.statusText}`;
  }
}
