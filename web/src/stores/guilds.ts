import { defineStore } from "pinia";

import type { Member } from "@/stores/user";
import http from "@/http";

export type Guild = {
  id: string;
  name: string;
  icon: string;
  owner: boolean;
  permissions: number;
  members?: Member[];
};

export const useGuildStore = defineStore("guild", {
  state: () => {
    return {
      guilds: [] as Guild[],
      loading: true,
      error: false,
    };
  },
  getters: {
    getGuildByID: (state) => {
      return (guildID: string): Guild | undefined =>
        state.guilds.find((guild) => guild.id === guildID);
    },
  },
  actions: {
    clear() {
      this.guilds = [];
    },
    async fetchGuilds() {
      const token = localStorage.getItem("token");
      this.loading = true;
      try {
        if (!token) {
          return;
        }
        const res = await http.get<Guild[]>("/api/guilds/", {
          headers: { Authorization: "Bearer " + token },
        });
        this.guilds = res.data;
        this.loading = false;
      } catch (error) {
        this.error = true;
        console.log(error);
      }
    },
    async fetchGuildMembers(guildID: string) {
      const g = this.getGuildByID(guildID);
      if (!g) {
        return;
      }

      try {
        const res = await http.get<Member[]>(
          "/api/guilds/" + guildID + "/members"
        );
        g.members = res.data;
      } catch (error) {
        console.log(error);
      }
    },
  },
});
