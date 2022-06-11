import { defineStore } from "pinia";

export const useNavStore = defineStore("nav", {
  state: () => {
    return {
      active: false,
    };
  },
  actions: {
    setActive(newState: boolean) {
      this.active = newState;
    },
  },
});
