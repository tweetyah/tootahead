export class Post{
  id?: number
  text?: string
  parentId?: string
  key?: string
  sendAt?: Date
  retweetAt?: Date
  threadCount?: number

  constructor(key?: string) {
    this.key = key
  }

  static getHtml(tweet: Post) {
    return tweet.text?.replace(/(?:\r\n|\r|\n)/g, '<br>')
  }

  static fromDb(row: any): Post {
    let tweet = new Post()
    tweet.id = row.id
    tweet.text = row.text
    tweet.parentId = row.tweet_parent
    tweet.threadCount = row.thread_count
    // TODO: Convert UTC to local time here
    if(row.send_at) tweet.sendAt = new Date(row.send_at)
    if(row.retweet_at) tweet.retweetAt = new Date(row.retweet_at)
    return tweet
  }
}

export type Auth = {
  access_token?: string
  id?: string
  name?: string
  profile_image_url?: string
  username?: string
}

export type TimeRange = {
  start: number
  end: number
}

export type Category = {
  id: number
  name: string
  color: string
}

export enum Service {
  Twitter = 0,
  Mastodon = 1
}

export type Alert = {
  title: string
  body: string
}