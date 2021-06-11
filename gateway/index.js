const { ApolloServer } = require('apollo-server');
const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
	willSendRequest({ request, context }) {
		request.http.headers.set('x-trace-id', context.xTraceId);
	}
}

const gateway = new ApolloGateway({
	serviceList: [
		{ name: 'user', url: 'http://localhost:8081/query' },
		{ name: 'role', url: 'http://localhost:8082/query' }, // Add here
	],
	buildService({ name, url }) {
		return new AuthenticatedDataSource({ url });
	},
	debug: true,
});

const server = new ApolloServer({
	gateway,
	subscriptions: false,
});

server.listen(8080).then(({ url }) => {
	console.log(`Server ready at ${url}`);
});
