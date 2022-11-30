import type { Post } from "./models";

export function SortPosts(posts: Post[], descending?: boolean) {
  if(descending) {
    posts.sort((a, b) => a.sendAt && b.sendAt ? b.sendAt.getTime() - a.sendAt.getTime() : 0)
  } else {
    posts.sort((a, b) => a.sendAt && b.sendAt ? a.sendAt.getTime() - b.sendAt.getTime() : 0)
  }
}