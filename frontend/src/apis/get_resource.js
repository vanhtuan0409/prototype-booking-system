function randInt() {
  return Math.round(Math.random() * 100);
}

export default function() {
  const resources = [];
  for (let i = 0; i < 36; i++) {
    resources.push({
      id: i + 1,
      available: randInt() % 2 === 0
    });
  }
  return Promise.resolve(resources);
}
