import { create } from 'zustand'
import { client } from '../api/client'

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
  loading: boolean;
  fetchChats: () => Promise<void>;
  fetchMessages: (chatId: number) => Promise<void>;
  setActiveChat: (id: number) => void;
  addMessage: (msg: Message) => void;
}

export const useChatStore = create<ChatState>((set, get) => ({
  chats: [],
  activeChatId: null,
  messages: [],
  loading: false,
  fetchChats: async () => {
    set({ loading: true });
    try {
      const res = await client.get('/chats');
      set({ chats: res.data, loading: false });
    } catch (err) {
      console.error('Failed to fetch chats', err);
      set({ loading: false });
    }
  },
  fetchMessages: async (chatId) => {
    try {
      const res = await client.get(`/chats/${chatId}/messages`);
      set({ messages: res.data });
    } catch (err) {
      console.error('Failed to fetch messages', err);
    }
  },
  setActiveChat: (id) => {
    set({ activeChatId: id });
    get().fetchMessages(id);
  },
  addMessage: (msg) => set((state) => {
    // Check if message is already in list to avoid duplicates from WS
    if (state.messages.some(m => m.id === msg.id)) return state;
    
    // Only add to view if it belongs to current chat
    if (msg.chatId === state.activeChatId) {
      return { messages: [...state.messages, msg] };
    }
    return state;
  })
}))
