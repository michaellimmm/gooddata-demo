import { emailRegex } from '../../utils/email';

export class LoginRequest {
  email: string = '';
  password: string = '';

  constructor(initializer?: any) {
    if (!initializer) return;

    if (initializer.email && typeof initializer.email === 'string') {
      this.email = initializer.email;
    }

    if (initializer.password && typeof initializer.password === 'string') {
      this.password = initializer.password;
    }
  }

  // return error message
  validate(): string[] {
    const errorMessages = Array<string>();
    if (this.email === '') {
      errorMessages.push('email can not be empty');
    } else if (!this.email.match(emailRegex)) {
      errorMessages.push('email is invalid');
    }

    if (this.password === '') {
      errorMessages.push('password can not be empty');
    }

    return errorMessages;
  }
}
