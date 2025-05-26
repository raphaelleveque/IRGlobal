// Configuração centralizada da aplicação
export const appConfig = {
  // API Configuration
  api: {
    baseUrl: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080",
    timeout: parseInt(import.meta.env.VITE_API_TIMEOUT || "10000"),
    endpoints: {
      auth: {
        login: "/auth/login",
        register: "/auth/register",
      },
      realizedPnl: "/realized-pnl/get",
      positions: "/position/get",
      transactions: {
        get: "/transaction/get",
        add: "/transaction/add",
      },
    },
  },

  // App Configuration
  app: {
    name: "IRGlobal",
    version: "1.0.0",
  },

  // Storage Keys
  storage: {
    authToken: "authToken",
    refreshToken: "refreshToken",
  },

  // UI Configuration
  ui: {
    defaultPageSize: 20,
    maxPageSize: 100,
    debounceDelay: 300,
    toastDuration: 5000,
  },

  // Currency Configuration
  currency: {
    default: "BRL",
    supported: ["BRL", "USD"],
    locale: "pt-BR",
  },
} as const;

// API URL builder
export const buildApiUrl = (endpoint: string) => {
  return `${appConfig.api.baseUrl}${endpoint}`;
};

// Get full endpoint URL
export const getEndpointUrl = (
  category: keyof typeof appConfig.api.endpoints,
  endpoint?: string
) => {
  const baseEndpoint = appConfig.api.endpoints[category];

  if (typeof baseEndpoint === "string") {
    return buildApiUrl(baseEndpoint);
  }

  if (
    endpoint &&
    typeof baseEndpoint === "object" &&
    endpoint in baseEndpoint
  ) {
    return buildApiUrl(baseEndpoint[endpoint as keyof typeof baseEndpoint]);
  }

  throw new Error(`Endpoint not found: ${category}.${endpoint}`);
};
