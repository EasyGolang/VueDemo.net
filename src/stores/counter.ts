export const useCounterStore = {
  id: 'counter',
  state: () => ({
    counter: 0,
  }),
  getters: {
    doubleCount: (state) => state.counter * 2,
  },
  actions: {
    increment() {},
  },
};
