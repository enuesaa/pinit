export type Binder = {
  id: number;
  name: string;
}

export type Note = {
  id: number;
  binderId: number;
  content: string;
}