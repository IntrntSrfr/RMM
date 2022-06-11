<template>
  <div class="member-wrapper">
    <div class="options">
      <div class="grp">
        <AppButton text="Select all" @click="markAll" />
        <AppButton text="Select none" @click="markNone" />
      </div>
      <div class="grp">
        <AppButton text="Ban selected" variant="danger" @click="banSelected" />
      </div>
    </div>
    <div class="member-list">
      <div
        class="member"
        :class="{ selected: m.checked }"
        v-for="(m, i) in members"
        :key="i"
        @click="mark(m)"
      >
        <div class="name">
          {{ `${m.member.user.username}#${m.member.user.discriminator}` }}
        </div>
        <div class="joined">{{ timeAgo.format(m.joined) }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Member } from "@/stores/user";
import AppButton from "@/components/AppButton.vue";
import TimeAgo from "javascript-time-ago";
import { useGuildStore } from "@/stores/guilds";

type TableMember = {
  member: Member;
  checked: boolean;
  joined: Date;
};

export default defineComponent({
  name: "GuildMembers",
  components: { AppButton },
  setup() {
    const guildStore = useGuildStore();
    const timeAgo = new TimeAgo("en-US");
    return { guildStore, timeAgo };
  },
  data() {
    return {
      members: [] as TableMember[],
    };
  },
  computed: {
    selectedMembers(): TableMember[] {
      return this.members.filter((m: TableMember) => m.checked);
    },
    guildMembers(): TableMember[] {
      const guildID = this.$route.params.guildID;
      const g = this.guildStore.getGuildByID(guildID.toString());
      if (!g || !g.members) {
        return [];
      }

      return g.members
        .map((m: Member) => {
          return {
            member: m,
            checked: false,
            joined: new Date(m.joined_at),
          } as TableMember;
        })
        .sort((a: TableMember, b: TableMember) => {
          return b.joined.getTime() - a.joined.getTime();
        });
    },
  },
  methods: {
    markAll() {
      this.members.forEach((m: TableMember) => (m.checked = true));
    },
    markNone(): void {
      this.members.forEach((m: TableMember) => (m.checked = false));
    },
    mark(tm: TableMember){
      tm.checked = !tm.checked;
    },
    banSelected() {
      if (!this.selectedMembers) {
        return;
      }
      const ans = confirm("Are you SURE you want to ban ALL selected users?");
      if (!ans) {
        return;
      }

      console.log(this.selectedMembers);
      console.log(this.selectedMembers.map((m) => m.member.user.id));
    },
  },
  mounted(){
    const guildID = this.$route.params.guildID;
    const g = this.guildStore.getGuildByID(guildID.toString());
    console.log(guildID, g);
    
    
    if (!g || !g.members) {
      return [];
    }

    const weed = g.members
      .map((m: Member) => {
        return {
          member: m,
          checked: false,
          joined: new Date(m.joined_at),
        } as TableMember;
      })
      .sort((a: TableMember, b: TableMember) => {
        return b.joined.getTime() - a.joined.getTime();
      });

    console.log(weed);
  },
});
</script>

<style scoped>
.options {
  display: flex;
  justify-content: space-between;
  gap: 1em;
  padding-bottom: 1em;
  border-bottom: 1px solid #666;
}

.grp{
  display:flex;
  gap:1em;
}

.member-list {
  display:flex;
  flex-direction:column;
  width: 100%;
  gap: 0.5em;
}

.member {

}

.name {
  font-size: 1rem;
  color: rgb(175, 175, 175);
}
.joined {
  font-size: 0.75rem;
  color: #666;
}
</style>
