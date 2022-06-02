import { defineStore } from "pinia";
import axios from "axios";

import { useUserStore, type Member } from "@/stores/user";
import http from "@/http";

export type Guild = {
  id: string;
  name: string;
  icon: string;
  owner: boolean;
  permissions: number;
  members?: Member[];
};

type GuildState = {
  guilds: Guild[];
};

export const useGuildStore = defineStore("guild", {
  state: () => {
    return {
      guilds: [],
    } as GuildState;
  },
  getters: {
    getGuildByID: (state) => {
      return (guildID: string): Guild | undefined =>
        state.guilds.find((guild) => guild.id === guildID);
    },
  },
  actions: {
    async fetchGuilds() {
      try {
        const token = useUserStore().token;
        if (!token?.access_token) {
          return;
        }
        const res = await axios.get<Guild[]>(
          "https://discord.com/api/users/@me/guilds",
          {
            headers: { Authorization: "Bearer " + token.access_token },
          }
        );
        this.guilds = res.data;
      } catch (error) {
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