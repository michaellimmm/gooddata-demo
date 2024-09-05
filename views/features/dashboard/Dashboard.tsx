import { useEffect, useState } from 'react';
import {
  IAuthenticationContext,
  NotAuthenticated,
} from '@gooddata/sdk-backend-spi';
import tigerFactory, {
  SetJwtCallback,
  TigerJwtAuthProvider,
} from '@gooddata/sdk-backend-tiger';
import { BackendProvider, WorkspaceProvider } from '@gooddata/sdk-ui';
import { idRef } from '@gooddata/sdk-model';
import { Dashboard } from '@gooddata/sdk-ui-dashboard';
import { useNavigate } from 'react-router-dom';
import { debounce } from 'lodash';
import { client } from '../../client';
import { Account } from '../../models/account';

const notAuthenticatedHandler = debounce(
  (_context: IAuthenticationContext, _error: NotAuthenticated) => {
    // handle the NotAuthenticated exception
    // Show error message on alert
  },
  500
);
const jwtIsAboutToExpire = async (setJwt: SetJwtCallback) => {
  const accessTokenJson = localStorage.getItem('account');
  if (accessTokenJson) {
    const account = JSON.parse(accessTokenJson);
    // tenant id
    client
      .AnalyticService_GetToken({
        tenantId: account['tenant_id'],
      })
      .then((result) => {
        const data = {
          name: result['name'] ?? '',
          email: result['email'] ?? '',
          tenantId: result['tenant_id'] ?? '',
          accessToken: result['access_token'] ?? '',
        } as Account;

        localStorage.setItem('account', JSON.stringify(data));

        // set the JWT back into authentication provider via provided callback
        setJwt(data.accessToken);
      })
      .catch((e) => {
        console.log(e);
      });
  }
};
const secondsBeforeTokenExpirationToCallReminder = 60;

export default function JourneyDashboard() {
  const [accessToken, setAccessToken] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const accessTokenJson = localStorage.getItem('account');
    if (accessTokenJson) {
      const { accessToken } = JSON.parse(accessTokenJson);
      setAccessToken(accessToken);
    } else {
      navigate('/', { replace: true });
    }
  }, []);

  if (accessToken === '') {
    return (
      <>
        <h1>Loading...</h1>
      </>
    );
  }

  return (
    <>
      <GoodDataBackendProvide accessToken={accessToken} />
    </>
  );
}

function GoodDataBackendProvide({ accessToken }: { accessToken: string }) {
  const backend = tigerFactory().withAuthentication(
    new TigerJwtAuthProvider(
      accessToken,
      notAuthenticatedHandler,
      jwtIsAboutToExpire,
      secondsBeforeTokenExpirationToCallReminder
    )
  );
  return (
    <BackendProvider backend={backend}>
      <WorkspaceProvider workspace={import.meta.env.VITE_WORKSPACEID}>
        <GoodDataDashboard />
      </WorkspaceProvider>
    </BackendProvider>
  );
}

const dashboard = idRef(import.meta.env.VITE_DASHBOARDID);

function GoodDataDashboard() {
  const config = {
    isReadOnly: true,
    hideSaveAsNewButton: true,
    isEmbedded: true,
    isWhiteLabeled: true,
  };

  return (
    <>
      <Dashboard dashboard={dashboard} config={config} />
    </>
  );
}
