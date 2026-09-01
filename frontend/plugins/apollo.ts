import { ApolloClient, InMemoryCache } from '@apollo/client/core';
import { provideApolloClient } from '@vue/apollo-composable';
import { useRuntimeConfig } from '#imports';

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig();
  const apolloClient = new ApolloClient({
    uri: `${config.public.hasuraGraphql}/v1/graphql`,
    cache: new InMemoryCache(),
  });

  provideApolloClient(apolloClient);
});
