export class Tweet{
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

  static getHtml(tweet: Tweet) {
    return tweet.text?.replace(/(?:\r\n|\r|\n)/g, '<br>')
  }

  static fromDb(row: any): Tweet {
    let tweet = new Tweet()
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
