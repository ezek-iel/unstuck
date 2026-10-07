export interface Question {
    id: number
    title: string
    details: string
    upvotes: number
    tags: string
}

export interface Comment {
    id: number
    details: string
    question: number
}


export async function fetchAllQuestions() {
    const request = new Request("http://localhost:1323/questions")
    const response = await fetch(request)
    return await response.json() as Question[]
}

export async function fetchQuestionDetails(questionId: number) {
    const response  = await fetch(`http://localhost:1323/questions/${questionId}`)
    return await response.json() as Question
}

export async function getQuestionComments(questionId: number) {
    const response = await fetch(`http://localhost:1323/questions/${questionId}/comments`)
    return await response.json() as Comment[]
}