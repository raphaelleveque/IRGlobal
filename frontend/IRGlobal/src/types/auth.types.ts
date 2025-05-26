export interface LoginCredentials {
  email: string;
  password: string;
}

export interface LoginFormData extends LoginCredentials {}

export interface AuthUser {
  id: string;
  email: string;
  name: string;
}

export interface LoginResponse {
  user: AuthUser;
  token: string;
}

export interface AuthError {
  message: string;
  code?: string;
}
