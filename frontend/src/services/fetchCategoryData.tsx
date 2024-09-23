/**
 * fetching all categories data
 */

export async function fetchCategoryData(categoryNum: string) {
  const apiUrl = "http://localhost:8080"

  // Change category_num type to int
  const categoryNumber = Number(categoryNum)

  const response = await fetch(`${apiUrl}/categories`);
  if (!response.ok) {
    throw new Error('Failed to fetch data');
  }
  const data = await response.json();
  
  // compare enter categoryNum and category.id in Database
  if (categoryNumber < data.length) {
    return data[categoryNumber-1].category_name;
  } else {
    console.log("Category Not Found.")
    return null;
  }
}