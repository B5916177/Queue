import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn';
import WelcomePage from './components/WelcomePage';
import QueueOrderCreate from './components/QueueOrder';
import Menu from './components/Menu';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/menu', Menu);
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/queueorder', QueueOrderCreate);
    
  }
});
