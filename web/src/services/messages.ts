import { reactive } from "vue";

const messages = reactive({
  data: [] as string[],
});

declare global {
  interface Window {
    SendMessage: (inp: string) => void;
  }
}

window.SendMessage = (inp: string) => {
  messages.data = [...messages.data, inp];
};

function useMessages() {
  return { messages };
}

export default useMessages;
