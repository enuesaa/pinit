import { mutate, type MutateConfig } from './base-mutate'
import { query } from './base-query'

export const queryGet = <T>(path: string) => query<T>('GET', path, {})

export const mutatePost = <R, T>(path: string, config: MutateConfig) => mutate<R, T>('POST', path, config)

export const mutatePut = <R, T>(path: string, config: MutateConfig) => mutate<R, T>('PUT', path, config)

export const mutateDelete = <R, T>(path: string, config: MutateConfig) => mutate<R, T>('DELETE', path, config)
