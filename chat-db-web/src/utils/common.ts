export function isJSON(str: any): boolean {
  try {
    JSON.parse(str);
    return true;
  } catch (e) {
    return false;
  }
}