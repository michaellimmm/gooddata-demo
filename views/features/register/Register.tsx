import { useState } from 'react';
import RegisterForm from './RegisterForm';
import { RegisterRequest } from './registerRequest';
import { useNavigate } from 'react-router-dom';
import { client } from '../../client';
import { Account } from '../../models/account';

export default function Register() {
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const onSubmit = (registerRequest: RegisterRequest) => {
    client
      .AnalyticService_RegisterAccount({
        email: registerRequest.email,
        name: registerRequest.name,
        tenant_id: registerRequest.tenantId,
        password: registerRequest.password,
      })
      .then((result) => {
        setErrorMessage('');

        const data = {
          name: result['name'] ?? '',
          email: result['email'] ?? '',
          tenantId: result['tenant_id'] ?? '',
          accessToken: result['access_token'] ?? '',
        } as Account;

        localStorage.setItem('account', JSON.stringify(data));
        navigate('/dashboard', { replace: true });
      })
      .catch((e) => {
        setErrorMessage(e.response.data.errors);
      });
  };

  return <RegisterForm errorMessage={errorMessage} onSubmit={onSubmit} />;
}
