import { Post } from "../models";
import type { Category } from "../models";

type Request = {
  method: string
  headers: {[key: string]: string}
  body?: string
}

export class ApiService {
  _base = "/.netlify/functions"
  _token

  constructor(token: string) {
    this._token = token
  }

  async fetchPosts(): Promise<Post[]> {
    const res = await this.execute("get", "/posts")
    let posts: Post[] = []
    res.forEach((r: any) => posts.push(Post.fromDb(r)))
    return posts
  }

  async savePosts(posts: Post[]) {
    return await this.execute("post", "/posts", JSON.stringify(posts))
  }

  async fetchCategories(): Promise<Category[]> {
    return await this.execute("get", "/categories")
  }

  async createCategory(name: string, color?: string): Promise<Category> {
    return await this.execute("post", "/categories", JSON.stringify({
      name,
      color
    }))
  }

  async execute(method: string, path: string, body?: string) {
    let req: Request = {
      method,
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${this._token}`
      }
    }
    if(body) {
      req.body = body
    }
    let res = await fetch(`${this._base}${path}`, req)
    return await res.json()
  }
}