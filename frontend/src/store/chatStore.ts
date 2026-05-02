import { create } from 'zustand'

export interface Message {
  id: number;
  chatId: number;
  sender: 'agent' | 'client';
  text: string;
  createdAt: string;
}

export interface Contact {
  id: number;
  name: string;
  phone: string;
  provider: string;
}

export interface Chat {
  id: number;
  contact: Contact;
  status: string;
}

interface ChatState {
  chats: Chat[];
  activeChatId: number | null;
  messages: Message[];
  setActiveChat: (id: number) => void;
  addMessage: (msg: Message) => void;
}

export const useChatStore = create<ChatState>((set) => ({
  chats: [
    { id: 1, contact: { id: 1, name: 'Иван Иванов', phone: '+79991234567', provider: 'telegram' }, status: 'open' },
    { id: 2, contact: { id: 2, name: 'ООО Ромашка', phone: '+79990001122', provider: 'whatsapp' }, status: 'open' }
  ],
  activeChatId: null,
  messages: [
    { id: 1, chatId: 1, sender: 'client', text: 'Здравствуйте, есть вопрос по API', createdAt: new Date().toISOString() },
    { id: 2, chatId: 1, sender: 'agent', text: 'Добрый день! Слушаю вас', createdAt: new Date().toISOString() }
  ],
  setActiveChat: (id) => set({ activeChatId: id }),
  addMessage: (msg) => set((state) => ({ messages: [...state.messages, msg] }))
}))
