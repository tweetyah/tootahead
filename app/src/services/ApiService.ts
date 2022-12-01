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

  async fetchPosts(filter?: string): Promise<Post[]> {
    let path = "/posts"
    if(filter) {
      path += `?filter=${filter}`
    }
    const res = await this.execute("get", path)
    if (res) {
      let posts: Post[] = []
      res.forEach((r: any) => posts.push(Post.fromDb(r)))
      return posts
    }
    return null
  }

  async fetchScheduledPosts(): Promise<Post[]> {
    const res = await this.execute("get", "/posts?filter=scheduled")
    let posts: Post[] = []
    res.forEach((r: any) => posts.push(Post.fromDb(r)))
    return posts
  }

  async savePosts(posts: Post[]) {
    return await this.execute("post", "/posts", JSON.stringify(posts))
  }

  async updatePosts(posts: Post[]) {
    return await this.execute("put", `/posts`, JSON.stringify(posts), true)
  }

  async deletePosts(posts: Post[]) {
    return await this.execute("delete", `/posts`, JSON.stringify(posts), true)
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

  async execute(method: string, path: string, body?: string, ignoreResponseBody?: boolean) {
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
    if(res.status < 200 || 299 < res.status ) {
      throw new Error("(ApiService.execute) non-success status code", res)
    }
    if(!ignoreResponseBody) {
      return await res.json()
    }
  }
}