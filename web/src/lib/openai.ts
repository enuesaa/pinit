import { OpenAI } from 'openai'
import { useMutation } from 'react-query'

export const useChat = (apiKey: string) =>
  useMutation({
    mutationFn: async (message: string): Promise<string> => {
      const openai = new OpenAI({
        apiKey,
        dangerouslyAllowBrowser: true, // TODO
      })

      const chatCompletion = await openai.chat.completions.create({
        messages: [{ role: 'user', content: message }],
        model: 'gpt-3.5-turbo',
      })

      console.log(chatCompletion)

      return chatCompletion.choices.at(0)?.message.content ?? ''
    },
  })
