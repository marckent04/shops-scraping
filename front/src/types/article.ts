export type Article = {
    shop: string;
    title: string;
    price: number;
    currency: string;
    image: string;
    detailsUrl: string;
}

export type SavedArticle = Article & {
    id: string;
    createdAt: Date
}