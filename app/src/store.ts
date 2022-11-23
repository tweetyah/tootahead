import { writable } from 'svelte/store';
import { Service, type Alert } from './models';
import type { ApiService } from './services/ApiService';

export const service = writable<Service>(Service.Twitter)
export const alert = writable<Alert>()
export const auth = writable<any>();
export const name = writable<string>();
export const profileImgUrl = writable<string>();
export const handle = writable<string>();

export const api = writable<ApiService>();

// TODO: Type this
export function setAuth(value: any) {
  auth.set(value)
  name.set(value.name)
  handle.set(value.username)
  profileImgUrl.set(value.profile_image_url)
}