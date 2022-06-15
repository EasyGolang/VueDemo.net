<script setup lang="ts">
import { nextTick } from 'vue';
import { cloneDeep, setToken } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { login } from '@/api/Account';

const $router = useRouter();

const from = $ref({
  Email: '',
  Password: '',
  resMsg: {},
});

const Submit = async (e: any) => {
  const Elm = e.target;
  Elm.disabled = true;
  const res = await login({
    Email: from.Email,
    Password: from.Password,
  });
  Elm.disabled = false;
  from.resMsg = res;

  if (res.Code > 0) {
    await setToken(res.Data.Token);
    $router.push('/about');
  }

  nextTick(() => {
    console.info(e);
    console.info(cloneDeep(from));
  });
};
</script>

<template>
  <main>
    <h1>登录</h1>
    <n-input
      type="text"
      size="small"
      v-model:value="from.Email"
      name="Email"
      placeholder="请输入邮箱"
      :inputProps="{ autocomplete: 'password' }"
    />
    <br />
    <br />
    <n-input
      type="password"
      size="small"
      :inputProps="{ autocomplete: 'password' }"
      v-model:value="from.Password"
      name="Password"
      placeholder="请输入密码"
    />
    <br />
    <br />
    Email: {{ from.Email }}
    <br />
    <br />
    Password: {{ from.Password }}
    <br />
    <br />

    <n-button type="primary" @click="Submit"> 提交 </n-button>
    <br />
    <br />

    返回信息：
    <div v-if="from.resMsg.Code" class="resMsg">{{ JSON.stringify(from.resMsg) }}</div>
  </main>
</template>

<style lang="less" scoped>
.resMsg {
  margin: 12px;
  border: 10px solid gray;
  padding: 15px;
  &.success {
    border-color: green;
  }
  &.fail {
    border-color: red;
  }
}
</style>
