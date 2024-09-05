import { useState } from 'react';
import { RegisterRequest } from './registerRequest';
import { Alert, Box, Button, TextField } from '@mui/material';

interface RegisterFormProps {
  errorMessage: string;
  onSubmit: (registerRequest: RegisterRequest) => void;
}

export default function RegisterForm({
  errorMessage,
  onSubmit,
}: RegisterFormProps) {
  const [error, setError] = useState<string[]>([]);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = new FormData(event.currentTarget);
    const registerRequest = new RegisterRequest({
      name: data.get('name'),
      email: data.get('email'),
      tenantId: data.get('tenantId'),
      password: data.get('password'),
    });
    const fieldErrorMessages = registerRequest.validate();
    if (fieldErrorMessages.length === 0) {
      onSubmit(registerRequest);
    }

    setError(fieldErrorMessages);
  };

  const fieldErrorMessages = error.map((messages) => {
    return (
      <Box mb={1}>
        <Alert severity="error">{messages}</Alert>
      </Box>
    );
  });

  return (
    <>
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
          {fieldErrorMessages}

          {errorMessage.length > 0 && (
            <Box mb={1}>
              <Alert severity="error">{errorMessage}</Alert>
            </Box>
          )}
          <TextField
            margin="normal"
            required
            fullWidth
            id="name"
            label="Name"
            name="name"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="tenantId"
            label="Tenant ID"
            name="tenantId"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Register
          </Button>
        </Box>
      </Box>
    </>
  );
}
