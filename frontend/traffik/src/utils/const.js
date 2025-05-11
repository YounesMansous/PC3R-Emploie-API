const BASE_URL = "http://localhost:8080";

const formatDate = (isoDate) => {
  const dateObj = new Date(isoDate);
  return dateObj.toLocaleString("fr-FR", {
    day: "2-digit",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

export { BASE_URL, formatDate };
