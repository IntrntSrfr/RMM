import { defineStore } from "pinia";
import axios from "axios";

export type UserState = {
  token: AccessTokenResponse | null;
  loggedIn: boolean;
  user: User | null;
  guilds: Guild[];
};

export type Member = {
  user: User;
  joinedAt: Date;
};

export type User = {
  id: string;
  username: string;
  discriminator: string;
  avatar: string;
  banner?: string;
  accent_color?: number;
};

export type AccessTokenResponse = {
  access_token: string;
  token_type: string;
  expires_in: number;
  refresh_token: string;
  scope: string;
};

export type Guild = {
  id: string;
  name: string;
  icon: string;
  owner: boolean;
  permissions: number;
};

export const useUserStore = defineStore("user", {
  state: () => {
    const t = localStorage.getItem("token");
    let tObj = null;
    if (t) {
      // IMPORTANT: CHECK EXPIRY DATE :DDD
      tObj = JSON.parse(t);
    }
    return {
      token: tObj,
      loggedIn: !!tObj,
      user: null,
      guilds: [],
    } as UserState;
  },
  getters: {
    getGuildByID: (state) => {
      return (guildID: string) =>
        state.guilds.find((guild) => guild.id === guildID);
    },
  },
  actions: {
    async fetchGuilds() {
      try {
        const token = this.token?.access_token;
        const res = await axios.get<Guild[]>(
          "https://discord.com/api/users/@me/guilds",
          {
            headers: { Authorization: "Bearer " + token },
          }
        );
        this.guilds = res.data;
      } catch (error) {
        console.log(error);
      }
    },
    async fetchUser() {
      try {
        const token = this.token?.access_token;
        const res = await axios.get<User>("https://discord.com/api/users/@me", {
          headers: { Authorization: "Bearer " + token },
        });
        console.log(res.data);
        this.loggedIn = true;
        this.user = res.data;
      } catch (error) {
        console.log(error);
      }
    },
    async oauth(code: string) {
      try {
        const res = await axios.get<AccessTokenResponse>(
          "http://localhost:4444/api/auth/callback",
          {
            params: { code: code },
          }
        );
        this.token = res.data;
        localStorage.setItem("token", JSON.stringify(res.data));
      } catch (e) {
        console.log(e);
      }
    },
  },
});
