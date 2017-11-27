package school

import (
	"sort"
)

type Grade struct {
 	grade int
 	students []string
 }

 type School struct {
 	grades []int
 	enrollment map[int]*Grade
 }
 
 func New() *School {
 	return &School{grades:make([]int, 0), enrollment: make(map[int]*Grade)}
 }
 
 func (s *School) Add(name string, grade int) {
	 g, ok := s.enrollment[grade]
 	if !ok {
		s.grades = append(s.grades, grade)
		g = &Grade{grade: grade}
		s.enrollment[grade] = g
	}
	 g.students = append(g.students, name)
	 sort.Strings(g.students)
 }
 
 func (s *School) Grade(grade int) []string {
 	var result []string
 	if g, ok := s.enrollment[grade]; ok {
 		result = g.students
	}
 	return result
 }
 
 func (s *School) Enrollment() []Grade {
 	result := make([]Grade, len(s.grades))
 	s.sort()
 	for i, g := range s.grades {
 		result[i] = *(s.enrollment[g])
	}
	return result
 }
 
 func (s *School) sort() {
 	sort.Ints(s.grades)
 }

