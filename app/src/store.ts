import { writable } from 'svelte/store';
import { Service, type Alert, type Auth, type CustomEmoji, type Instance } from './models';
import { ApiService } from './services/ApiService';

export const service = writable<Service>(Service.Twitter)
export const alert = writable<Alert>()
export const auth = writable<any>();
export const name = writable<string>();
export const profileImgUrl = writable<string>();
export const handle = writable<string>();
export const instance = writable<Instance>()
export const custom_emoji = writable<CustomEmoji[]>()

export const api = writable<ApiService>();

export async function init() {
  let authItem = localStorage.getItem("auth")
  if(authItem) {
    let authValue: Auth = JSON.parse(authItem)

    auth.set(authValue)
    name.set(authValue.name)
    handle.set(authValue.username)
    profileImgUrl.set(authValue.profile_image_url)

    api.set(new ApiService(authValue.access_token))
  }

  let instanceItem = localStorage.getItem("instance")
  if (instanceItem) {
    let instanceValue: Instance = JSON.parse(instanceItem)
    instance.set(instanceValue)

    let res = await fetch(`https://${instanceValue.domain}/api/v1/custom_emojis`)
    let json = await res.json()
    custom_emoji.set(json)
  }
}

export function setAuth(value: any) {
  auth.set(value)
  name.set(value.name)
  handle.set(value.username)
  profileImgUrl.set(value.profile_image_url)
}