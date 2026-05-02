import { useState } from 'react';
import { useChatStore } from './store/chatStore';
import { Search, Send, Paperclip, MoreVertical, Phone, Video } from 'lucide-react';

function App() {
  const { chats, activeChatId, messages, setActiveChat, addMessage } = useChatStore();
  const [inputText, setInputText] = useState('');

  const activeChat = chats.find(c => c.id === activeChatId);
  const activeMessages = messages.filter(m => m.chatId === activeChatId);

  const handleSend = () => {
    if (!inputText.trim() || !activeChatId) return;
    addMessage({
      id: Date.now(),
      chatId: activeChatId,
      sender: 'agent',
      text: inputText,
      createdAt: new Date().toISOString(),
    });
    setInputText('');
  };

  return (
    <div className="flex h-screen bg-[#0f172a] text-slate-200 font-sans overflow-hidden">
      {/* Sidebar - Chat List */}
      <div className="w-80 bg-[#1e293b] border-r border-slate-800 flex flex-col">
        <div className="p-4 border-b border-slate-800 flex items-center justify-between">
          <h1 className="text-xl font-semibold bg-gradient-to-r from-blue-400 to-indigo-500 bg-clip-text text-transparent">OmniCRM</h1>
        </div>
        <div className="p-3">
          <div className="relative">
            <Search className="absolute left-3 top-2.5 w-4 h-4 text-slate-400" />
            <input 
              type="text" 
              placeholder="Search chats..." 
              className="w-full bg-slate-800 text-sm text-slate-200 rounded-full pl-9 pr-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all border border-slate-700"
            />
          </div>
        </div>
        <div className="flex-1 overflow-y-auto">
          {chats.map(chat => (
            <div 
              key={chat.id} 
              onClick={() => setActiveChat(chat.id)}
              className={`flex items-center gap-3 p-4 cursor-pointer transition-colors border-b border-slate-800/50 ${activeChatId === chat.id ? 'bg-blue-500/10 border-l-4 border-l-blue-500' : 'hover:bg-slate-800 border-l-4 border-l-transparent'}`}
            >
              <div className="w-12 h-12 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-white font-bold shadow-lg shrink-0">
                {chat.contact.name.charAt(0)}
              </div>
              <div className="flex-1 min-w-0">
                <div className="flex justify-between items-baseline">
                  <h3 className="font-medium text-slate-100 truncate">{chat.contact.name}</h3>
                  <span className="text-xs text-slate-500 shrink-0">12:30</span>
                </div>
                <p className="text-sm text-slate-400 truncate flex items-center gap-1">
                  <span className={`w-2 h-2 rounded-full ${chat.contact.provider === 'whatsapp' ? 'bg-green-500' : 'bg-blue-400'}`}></span>
                  <span className="truncate">Last message preview...</span>
                </p>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Main Chat Area */}
      <div className="flex-1 flex flex-col bg-[#0b1120] relative">
        {activeChat ? (
          <>
            {/* Chat Header */}
            <div className="h-16 px-6 bg-[#1e293b]/80 backdrop-blur-md border-b border-slate-800 flex items-center justify-between z-10 sticky top-0">
              <div className="flex items-center gap-3">
                <div className="w-10 h-10 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-white font-bold shrink-0">
                  {activeChat.contact.name.charAt(0)}
                </div>
                <div>
                  <h2 className="font-semibold text-slate-100 leading-tight">{activeChat.contact.name}</h2>
                  <p className="text-xs text-slate-400 capitalize">{activeChat.contact.provider} • Online</p>
                </div>
              </div>
              <div className="flex items-center gap-4 text-slate-400">
                <Phone className="w-5 h-5 hover:text-slate-200 cursor-pointer transition-colors" />
                <Video className="w-5 h-5 hover:text-slate-200 cursor-pointer transition-colors" />
                <MoreVertical className="w-5 h-5 hover:text-slate-200 cursor-pointer transition-colors" />
              </div>
            </div>

            {/* Messages */}
            <div className="flex-1 overflow-y-auto p-6 space-y-6">
              {activeMessages.map((msg) => (
                <div key={msg.id} className={`flex ${msg.sender === 'agent' ? 'justify-end' : 'justify-start'}`}>
                  <div className={`max-w-[70%] p-4 rounded-2xl shadow-sm ${msg.sender === 'agent' ? 'bg-gradient-to-br from-blue-500 to-indigo-600 text-white rounded-br-none' : 'bg-slate-800 text-slate-200 rounded-bl-none border border-slate-700'}`}>
                    <p className="leading-relaxed whitespace-pre-wrap">{msg.text}</p>
                    <span className={`text-[10px] mt-2 block ${msg.sender === 'agent' ? 'text-blue-200' : 'text-slate-400'}`}>
                      {new Date(msg.createdAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </span>
                  </div>
                </div>
              ))}
            </div>

            {/* Input Area */}
            <div className="p-4 bg-[#1e293b] border-t border-slate-800">
              <div className="flex items-center gap-3 bg-slate-900 rounded-full p-2 pl-4 border border-slate-700 focus-within:border-blue-500 focus-within:ring-1 focus-within:ring-blue-500 transition-all shadow-inner">
                <Paperclip className="w-5 h-5 text-slate-400 cursor-pointer hover:text-slate-200" />
                <input
                  type="text"
                  value={inputText}
                  onChange={(e) => setInputText(e.target.value)}
                  onKeyDown={(e) => e.key === 'Enter' && handleSend()}
                  placeholder="Type a message..."
                  className="flex-1 bg-transparent border-none focus:outline-none text-slate-200 placeholder-slate-500 text-sm"
                />
                <button 
                  onClick={handleSend}
                  className="w-10 h-10 shrink-0 rounded-full bg-blue-500 hover:bg-blue-600 flex items-center justify-center text-white transition-transform hover:scale-105 active:scale-95 shadow-md disabled:opacity-50"
                  disabled={!inputText.trim()}
                >
                  <Send className="w-4 h-4 ml-0.5" />
                </button>
              </div>
            </div>
          </>
        ) : (
          <div className="flex-1 flex flex-col items-center justify-center text-slate-500 bg-gradient-to-b from-[#0b1120] to-[#0f172a]">
            <div className="w-24 h-24 mb-6 rounded-full bg-slate-800/50 flex items-center justify-center shadow-inner border border-slate-700/50">
              <Search className="w-10 h-10 text-slate-600" />
            </div>
            <p className="text-xl font-medium text-slate-300 mb-2">No Chat Selected</p>
            <p className="text-sm">Choose a conversation from the left menu</p>
          </div>
        )}
      </div>

      {/* Right Sidebar - CRM Info */}
      {activeChat && (
        <div className="w-80 bg-[#1e293b] border-l border-slate-800 flex flex-col transform transition-transform duration-300 ease-in-out hidden lg:flex">
          <div className="p-8 border-b border-slate-800 flex flex-col items-center bg-gradient-to-b from-slate-800/50 to-transparent">
            <div className="w-24 h-24 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-white font-bold text-4xl mb-4 shadow-lg ring-4 ring-slate-900/50">
              {activeChat.contact.name.charAt(0)}
            </div>
            <h2 className="text-xl font-semibold text-slate-100 text-center">{activeChat.contact.name}</h2>
            <p className="text-sm text-slate-400 mt-1 flex items-center gap-1">
              <Phone className="w-3 h-3" />
              {activeChat.contact.phone}
            </p>
          </div>
          
          <div className="p-6 flex-1 overflow-y-auto">
            <div className="space-y-6">
              <div>
                <h3 className="text-xs uppercase font-bold text-slate-500 tracking-wider mb-3">CRM Details</h3>
                <div className="space-y-3">
                  <div className="bg-slate-900/50 p-3.5 rounded-xl border border-slate-700/50 hover:border-slate-600 transition-colors">
                    <span className="text-xs text-slate-400 block mb-1">Status</span>
                    <span className="text-sm font-medium text-green-400 bg-green-400/10 px-2.5 py-1 rounded-md inline-block border border-green-400/20">Active Lead</span>
                  </div>
                  <div className="bg-slate-900/50 p-3.5 rounded-xl border border-slate-700/50 hover:border-slate-600 transition-colors">
                    <span className="text-xs text-slate-400 block mb-1">Channel</span>
                    <span className="text-sm text-slate-200 capitalize flex items-center gap-2">
                      <div className={`w-2 h-2 rounded-full ${activeChat.contact.provider === 'whatsapp' ? 'bg-green-500' : 'bg-blue-400'}`}></div>
                      {activeChat.contact.provider}
                    </span>
                  </div>
                </div>
              </div>
              
              <div>
                <h3 className="text-xs uppercase font-bold text-slate-500 tracking-wider mb-3 flex justify-between items-center">
                  Notes
                  <span className="text-[10px] bg-slate-800 px-2 py-0.5 rounded text-slate-400">Auto-save</span>
                </h3>
                <textarea 
                  className="w-full bg-slate-900/80 border border-slate-700/50 rounded-xl p-3.5 text-sm text-slate-300 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 resize-none h-32 transition-all placeholder-slate-600 shadow-inner"
                  placeholder="Add a note about this contact..."
                  defaultValue="Interested in enterprise API integration. Follow up on Tuesday."
                ></textarea>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
