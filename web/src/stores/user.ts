import { defineStore } from "pinia";
import axios from "axios";

import http from "@/http";
import jwtDecode from "jwt-decode";

export type UserState = {
  token: JWT | null;
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

export type JWT = {
  exp: number; // expiry time
  iat: number; // issued at
  iss: string; // issued by
  sub: string; // user id
  token: string; // access token
  token_type: string; // token type
  refresh_token: string; // refresh token
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
    const t = localStorage.getItem("token_data");
    let tObj: JWT | null = null;
    if (t) {
      // IMPORTANT: CHECK EXPIRY DATE :DDD
      tObj = JSON.parse(t);
    }
    return {
      token: tObj,
      loggedIn: false,
      user: null,
    } as UserState;
  },
  getters: {},
  actions: {
    async fetchUser() {
      try {
        const token = this.token?.token;
        if (!token) {
          return;
        }
        const res = await axios.get<User>("https://discord.com/api/users/@me", {
          headers: { Authorization: "Bearer " + token },
        });
        this.loggedIn = true;
        this.user = res.data;
      } catch (error) {
        this.token = null;
        this.loggedIn = false;
        this.user = null;
      }
    },
    async oauth(code: string) {
      type JwtResp = { token: string };

      try {
        const res = await http.get<JwtResp>("/api/auth/callback", {
          params: { code: code },
        });
        const decoded = jwtDecode<JWT>(res.data.token);
        this.token = decoded;
        localStorage.setItem("id", this.token.sub);
        localStorage.setItem("token", res.data.token);
        localStorage.setItem("token_data", JSON.stringify(this.token));
      } catch (e) {
        this.token = null;
      }
    },
  },
});
