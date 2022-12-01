export enum ViewState {
  None,
  Done,
  Loading,
  NoData
}

export class Post{
  id?: number
  text?: string
  parentId?: string
  key?: string
  sendAt?: Date
  resendAt?: Date
  threadCount?: number

  constructor(key?: string) {
    this.key = key
  }

  html() {
    return this.text?.replace(/(?:\r\n|\r|\n)/g, '<br>')
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
    if(row.sendAt) tweet.sendAt = new Date(row.sendAt)
    if(row.resendAt) tweet.resendAt = new Date(row.resendAt)
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

export interface Instance {
  domain: string
  title: string
  version: string
  source_url: string
  description: string
  usage: Usage
  thumbnail: Thumbnail
  languages: string[]
  configuration: Configuration
  registrations: Registrations
  contact: Contact
  rules: Rule[]
}

export interface Usage {
  users: Users
}

export interface Users {
  active_month: number
}

export interface Thumbnail {
  url: string
  blurhash: string
  versions: Versions
}

export interface Versions {
  "@1x": string
  "@2x": string
}

export interface Configuration {
  urls: Urls
  accounts: Accounts
  statuses: Statuses
  media_attachments: MediaAttachments
  polls: Polls
  translation: Translation
}

export interface Urls {
  streaming: string
}

export interface Accounts {
  max_featured_tags: number
}

export interface Statuses {
  max_characters: number
  max_media_attachments: number
  characters_reserved_per_url: number
}

export interface MediaAttachments {
  supported_mime_types: string[]
  image_size_limit: number
  image_matrix_limit: number
  video_size_limit: number
  video_frame_rate_limit: number
  video_matrix_limit: number
}

export interface Polls {
  max_options: number
  max_characters_per_option: number
  min_expiration: number
  max_expiration: number
}

export interface Translation {
  enabled: boolean
}

export interface Registrations {
  enabled: boolean
  approval_required: boolean
  message: any
}

export interface Contact {
  email: string
  account: Account
}

export interface Account {
  id: string
  username: string
  acct: string
  display_name: string
  locked: boolean
  bot: boolean
  discoverable: boolean
  group: boolean
  created_at: string
  note: string
  url: string
  avatar: string
  avatar_static: string
  header: string
  header_static: string
  followers_count: number
  following_count: number
  statuses_count: number
  last_status_at: string
  noindex: boolean
  emojis: any[]
  fields: Field[]
}

export interface Field {
  name: string
  value: string
  verified_at: any
}

export interface Rule {
  id: string
  text: string
}

export interface CustomEmoji{
  shortcode: string
  url: string
  static_url: string
  visible_in_picker: boolean
}

