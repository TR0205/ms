export const resolvers = {
    Query: {
        book: async (parent, args, context) => books[args.id-1],
        books: async (parent, args, context) => books,
    //   book: async (parent, args, context) => {
    //     const response = await context.dataSources.catalogueApi.getBook(args.id)
    //     return response
    //   },
    //   books: async (parent, args, context) => {
    //     const response = await context.dataSources.catalogueApi.listBooks()
    //     return response
    //   },
    }
}

const books = [
    {
        id: 1,
        title: "The Awaking",
        author: "Kate Chopin",
        price: 1000
    }
];