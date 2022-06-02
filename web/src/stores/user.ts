import { defineStore } from "pinia";
import axios from "axios";

import http from "@/http";

export type UserState = {
  token: AccessTokenResponse | null;
  loggedIn: boolean;
  user: User | null;
};

export type Member = {
  user: User;
  joined_at: string;
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
    } as UserState;
  },
  getters: {},
  actions: {
    async fetchUser() {
      try {
        const token = this.token;
        if (!token?.access_token) {
          return;
        }
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
        const res = await http.get<AccessTokenResponse>("/api/auth/callback", {
          params: { code: code },
        });
        this.token = res.data;
        localStorage.setItem("token", JSON.stringify(res.data));
      } catch (e) {
        console.log(e);
      }
    },
  },
});
