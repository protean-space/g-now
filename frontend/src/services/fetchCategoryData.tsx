export async function fetchCategoryData(categoryNum: string) {
  // Replace with your actual API endpoint
  const apiUrl = "http://localhost:8080"

  // Change category_num type to int
  const categoryNumber = Number(categoryNum)

  const response = await fetch(`${apiUrl}/categories`);
  console.log(`category: ${categoryNumber}`)
  console.dir(response, {depth:null})
  if (!response.ok) {
    throw new Error('Failed to fetch data');
  }
  const data = await response.json();
  console.log(data)
  
  // compare enter categoryNum and category.id in Database
  if (categoryNumber < data.length) {
    return data[categoryNumber-1].category_name;
  } else {
    console.log("Category Not Found.")
    return null;
  }
}