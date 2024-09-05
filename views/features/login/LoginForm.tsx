import { Alert, Box, Button, Grid, Link, TextField } from '@mui/material';
import { useState } from 'react';
import { LoginRequest } from './loginRequest';
import { Link as RouterLink } from 'react-router-dom';

interface LoginFormProps {
  errorMessage: string;
  onSubmit: (loginRequest: LoginRequest) => void;
}

export default function LoginForm({ onSubmit, errorMessage }: LoginFormProps) {
  const [error, setError] = useState<string[]>([]);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = new FormData(event.currentTarget);
    const loginRequest = new LoginRequest({
      email: data.get('email'),
      password: data.get('password'),
    });
    const fieldErrorMessages = loginRequest.validate();
    if (fieldErrorMessages.length === 0) {
      onSubmit(loginRequest);
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
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
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
            autoComplete="current-password"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
          <Grid container>
            <Grid item>
              <Link component={RouterLink} to="/register" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </>
  );
}
