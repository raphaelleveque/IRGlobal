export interface LoginCredentials {
  email: string;
  password: string;
}

export interface RegisterCredentials {
  name: string;
  email: string;
  password: string;
}

export interface LoginFormData extends LoginCredentials {}
export interface RegisterFormData extends RegisterCredentials {}

export interface AuthUser {
  id: string;
  email: string;
  name: string;
}

export interface LoginResponse {
  user: AuthUser;
  token: string;
}

export interface RegisterResponse {
  user: AuthUser;
  token: string;
}

export interface AuthError {
  message: string;
  code?: string;
}
