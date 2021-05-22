export type PostType = {
  slug: string;
  title: string;
  content: string;
  html: string;
};

export type PostListType = PostType[];

export type ArchiveItemType = {
  slug: string;
  title: string;
  created_at: string;
};

export type ArchiveItemsByYearType = {
  [year: number]: ArchiveItemType[];
};
