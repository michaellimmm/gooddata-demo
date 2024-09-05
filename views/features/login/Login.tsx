import { useEffect, useState } from 'react';
import { client } from '../../client';
import LoginForm from './LoginForm';
import { useNavigate } from 'react-router-dom';
import { Account } from '../../models/account';
import { LoginRequest } from './loginRequest';

export default function Login() {
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const accessToken = localStorage.getItem('account');

    if (accessToken !== null) {
      navigate('/dashboard', { replace: true });
    }
  }, []);

  const onSubmit = (loginRequest: LoginRequest) => {
    client
      .AnalyticService_Login({
        email: loginRequest.email,
        password: loginRequest.password,
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

  return <LoginForm onSubmit={onSubmit} errorMessage={errorMessage} />;
}
