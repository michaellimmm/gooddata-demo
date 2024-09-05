import { emailRegex } from '../../utils/email';

export class RegisterRequest {
  name: string = '';
  email: string = '';
  password: string = '';
  tenantId: string = '';

  constructor(initializer?: any) {
    if (!initializer) return;

    if (initializer.name && typeof initializer.name === 'string') {
      this.name = initializer.name;
    }

    if (initializer.email && typeof initializer.email === 'string') {
      this.email = initializer.email;
    }

    if (initializer.password && typeof initializer.password === 'string') {
      this.password = initializer.password;
    }

    if (initializer.tenantId && typeof initializer.tenantId === 'string') {
      this.tenantId = initializer.tenantId;
    }
  }

  validate(): string[] {
    const errorMessages = Array<string>();

    if (this.name === '') {
      errorMessages.push('name can not be empty');
    }

    if (this.email === '') {
      errorMessages.push('email can not be empty');
    } else if (!this.email.match(emailRegex)) {
      errorMessages.push('email is invalid');
    }

    if (this.password === '') {
      errorMessages.push('password can not be empty');
    }

    if (this.tenantId === '') {
      errorMessages.push('tenant_id can not be empty');
    }

    return errorMessages;
  }
}
